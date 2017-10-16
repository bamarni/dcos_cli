import click

from dcos_cli.cli.cluster.attach import attach
from dcos_cli.cli.cluster.list import cluster_list
from dcos_cli.cli.cluster.remove import remove
from dcos_cli.cli.cluster.rename import rename
from dcos_cli.cli.cluster.setup import setup


@click.group()
def cluster():
    """Manage your DC/OS clusters."""
    pass


cluster.add_command(attach)
cluster.add_command(cluster_list)
cluster.add_command(remove)
cluster.add_command(rename)
cluster.add_command(setup)
