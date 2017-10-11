import click

from dcos_cli.cli.cluster import cluster

@cluster.command()
def attach():
    """Attach to a cluster."""
    click.echo('Attach to a cluster.')
