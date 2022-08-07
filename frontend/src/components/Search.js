const Search = () => {
  return (
    <form action="/" method="get" id="search-box">
      <input
        type="text"
        placeholder="Search transaction"
        name="search" 
      />
      <button type="submit">Search</button>
    </form>
  )
}

export default Search;