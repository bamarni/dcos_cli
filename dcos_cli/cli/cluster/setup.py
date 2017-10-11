import click

from dcos_cli.cli.cluster import cluster

@cluster.command()
def setup():
    """Setup a cluster."""
    print('Setup a cluster.')
