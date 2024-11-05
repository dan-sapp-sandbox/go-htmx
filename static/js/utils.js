function updateQueryParams(pokemonId) {
  const url = new URL(window.location);
  const params = url.searchParams;
  const ids = params.get('id') ? params.get('id').split(',') : [];

  if (ids.includes(pokemonId)) {
    const index = ids.indexOf(pokemonId);
    if (index > -1) {
      ids.splice(index, 1);
    }
  } else {
    ids.push(pokemonId);
  }

  if (ids.length > 0 && ids.length < 7) {
    params.set('id', ids.join(','));
  } else if (ids.length > 0) {
    return
  } else {
    params.delete('id');
  }
  
  window.history.pushState({}, '', url);
}