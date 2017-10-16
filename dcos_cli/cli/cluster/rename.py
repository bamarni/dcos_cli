from dcos_cli.cli.cluster import cluster


@cluster.command()
def rename():
    """Rename a cluster."""
    print('Rename a cluster.')
