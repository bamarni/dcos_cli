from dcos_cli.cli.cluster import cluster


@cluster.command(name='list')
def cluster_list():
    """List clusters."""
    print('List clusters.')
