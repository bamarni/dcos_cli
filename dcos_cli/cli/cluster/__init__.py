from dcos_cli.cli.dcos import dcos


@dcos.group()
def cluster():
    """Manage your DC/OS clusters."""
    pass


# pylint: disable=wrong-import-position
import dcos_cli.cli.cluster.attach  # noqa: E402,F401
import dcos_cli.cli.cluster.list  # noqa: E402,F401
import dcos_cli.cli.cluster.remove  # noqa: E402,F401
import dcos_cli.cli.cluster.rename  # noqa: E402,F401
import dcos_cli.cli.cluster.setup  # noqa: E402,F401
