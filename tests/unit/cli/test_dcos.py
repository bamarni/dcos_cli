from click.testing import CliRunner

from dcos_cli import __version__
from dcos_cli.cli.dcos import dcos


def test_version():
    runner = CliRunner()
    result = runner.invoke(dcos, ['--version'])
    assert result.exit_code == 0
    assert result.output == __version__ + '\n'
