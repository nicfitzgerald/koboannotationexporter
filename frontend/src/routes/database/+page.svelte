<script>
  import { GetInfoFromDB } from '$lib/wailsjs/go/main/App'
  let dbInfo = $state('')

  let items = $state('')
  async function getInfoFromDB() {
    items = await GetInfoFromDB()
    return items
  }
</script>

<h1>Database Info!</h1>
<button onclick={getInfoFromDB}>Get Info</button>
<p>This is what the DB shows:</p>
{#each items.annotations as author}
  <br />
  <h1>{author.name}</h1>
  {#each author.books as book}
    <h2>{book.title}</h2>
    {#each book.excerpts as excerpt}
      <span>{excerpt.text}, created at {excerpt.createdAt}</span><br />
    {/each}
  {/each}
{/each}
