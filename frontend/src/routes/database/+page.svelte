<script>
  import { GetInfoFromDB, GetDBPath } from '$lib/wailsjs/go/main/App'
  let items = $state('')
  let path = $state('')

  function getDBPath() {
    GetDBPath().then((result) => {
      path = result
    })
  }
  function getInfoFromDB() {
    GetInfoFromDB(path).then((result) => {
      items = result
    })
    return items
  }
</script>

<h1>Database Info!</h1>
<button onclick={getDBPath}>Select DB</button><br />
<span>The current path is: {path}</span><br />
<br />
<hr />
<br />
<button onclick={getInfoFromDB}>Get Info</button>
<p>This is what the DB shows:</p>
{#each items.annotations as author}
  <br />
  <h1>{author.name}</h1>
  {#each author.books as book}
    <h2>{book.title}</h2>
    {#each book.excerpts as excerpt}
      <span>{excerpt.text} (created at <strong>{excerpt.createdAt}</strong>)</span><br />
    {/each}
  {/each}
{/each}