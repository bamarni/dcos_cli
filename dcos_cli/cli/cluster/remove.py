from dcos_cli.cli.cluster import cluster


@cluster.command()
def remove():
    """Remove cluster(s)."""
    print('Remove cluster(s).')
