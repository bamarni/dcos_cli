package config

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"sort"

	"github.com/dcos/dcos-cli/pkg/config/tomlutil"
	"github.com/pelletier/go-toml"
)

// Config for the DC/OS CLI. The "core" and "cluster" sections are mapped
// to a struct. Additional sections ("marathon", etc.) are stored in a Tree.
type Config struct {
	Core core `toml:"core"`
	Cluster cluster `toml:"cluster"`
	tree *toml.Tree
}

type core struct {
	ACSToken string `toml:"dcos_acs_token"`
	URL string `toml:"dcos_url"`
	SSL ssl `toml:"ssl_verify"`
}

type ssl string

func (s ssl) Insecure() bool {
	if s == "false" || s == "False" {
		return true
	}
	return false
}

func (s ssl) CAPath() string {
	if s.Insecure() {
		return ""
	}
	return string(s)
}

type cluster struct {
	Name string `toml:"name"`
}

func FromPath(path string) (Config, error) {
	tree, err := toml.LoadFile(path)
	if err != nil {
		return Config{}, err
	}
	return fromTree(tree)
}

func FromString(tomlData string) (Config, error) {
	tree, err := toml.Load(tomlData)
	if err != nil {
		return Config{}, err
	}
	return fromTree(tree)
}

func (cfg Config) String() string {
	var cfgStr string

	// ordering here is not stable
	sections := cfg.tree.Keys()

	// sort sections alphabetically, except for "core" and "cluster",
	// which should respectively be the first and second one
	sort.SliceStable(sections, func(i, j int) bool {
		for _, section := range []string{"core", "cluster"} {
			if sections[i] == section || sections[j] == section {
				return sections[i] == section
			}
		}
		return sections[i] < sections[j]
	})

	for _, section := range sections {
		if sectionStr, err := cfg.StringAt([]string{section}); err == nil {
			cfgStr += sectionStr
		}
	}
	return cfgStr
}

func (cfg Config) StringAt(key []string) (string, error) {
	val := cfg.get(key)

	switch value := val.(type) {
	case *toml.Tree:
		var buf bytes.Buffer
		// walk tree leaves and print their value prefixed by their key
		tomlutil.WalkLeaves(value, key, func (leaf interface{}, nodeKeys []string) {
			strVal := fmt.Sprintf("%v\n", leaf)
			strKey := strings.Join(nodeKeys, ".")
			buf.Write([]byte(fmt.Sprintf("%s %v", strKey, strVal)))
		})
		return buf.String(), nil
	default:
		// this is a leaf holding an interface value we can print,
		// there is no need to print the key before the value
		if value == nil {
			strKey := strings.Join(key, ".")
			return "", fmt.Errorf("unknown config key \"%s\"", strKey)
		}
		return fmt.Sprintf("%v\n", val), nil
	}
}

func (c core) String() string {
	cfgStr := fmt.Sprintf("%s %v\n", "core.dcos_acs_token", "********")
	cfgStr = cfgStr + fmt.Sprintf("%s %v\n", "core.dcos_url", c.URL)
	cfgStr = cfgStr + fmt.Sprintf("%s %v", "core.ssl_verify", c.SSL)

	return cfgStr
}

func (c cluster) String() string {
	cfgStr := fmt.Sprintf("%s %v", "cluster.name", c.Name)

	return cfgStr
}

func fromTree(tree *toml.Tree) (Config, error) {
	var cfg Config
	if err := tree.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	cfg.tree = tree

	return cfg, nil
}

func (cfg Config) get(key []string) interface{} {
	if len(key) == 0 {
		panic("no key specified")
	}
	keyPos := 0
	keyLen := len(key)

	t := reflect.TypeOf(cfg)
	v := reflect.ValueOf(cfg)
redo:
	for i := 0; t.Kind() == reflect.Struct && i < t.NumField(); i++ {
		field := t.Field(i)
		val, ok := field.Tag.Lookup("toml")
		if ok && val == key[keyPos] {
			t = field.Type
			v = v.FieldByName(field.Name)
			keyPos++

			if keyPos == keyLen {
				return v.Interface()
			}
			goto redo
		}
	}
	return cfg.tree.GetPath(key)
}