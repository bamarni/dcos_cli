import click

from dcos_cli.cli.cluster import cluster

@cluster.command()
def list():
    """List clusters."""
    print('List clusters.')
