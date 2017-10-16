import click


@click.command(name='list')
def cluster_list():
    """List clusters."""
    print('List clusters.')
