<script>
  import '../app.css'

  let { children } = $props()
</script>

<nav>
  <a href="/">home</a>
  <a href="/settings">settings</a>
  <a href="/database">database</a>
</nav>

{@render children()}
