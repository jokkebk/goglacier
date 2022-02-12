<script>
    import { onMount } from "svelte";

    let dbs = [];

    onMount(async function() {
        const response = await fetch('http://localhost:8080/dbs');
        dbs = await response.json();
    });
</script>
<h1>Welcome to GoGlacier</h1>
<p>Choose a backup DB:</p>

<ul>
    {#each dbs as db}
    <li>{db}</li>
    {/each}
</ul>