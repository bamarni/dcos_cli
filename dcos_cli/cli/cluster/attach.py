
from dcos_cli.cli.cluster import cluster


@cluster.command()
def attach():
    """Attach to a cluster."""
    print('Attach to a cluster.')
