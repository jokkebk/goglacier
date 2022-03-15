<script>
    import { page } from '$app/stores';
    import { onMount } from "svelte";

    let entry = [];
    let scans = [];
    let entryId = $page.params.id;
    let files = [];
    let folders = [];

    onMount(async function () {
        let resp = await fetch("http://localhost:8080/entries");
        let entries = await resp.json();
        entry = entries.find(e => e.Id == entryId);

        resp = await fetch("http://localhost:8080/scans");
        scans = await resp.json();
        scans = scans.filter(s => s.EntryId == entryId);

        resp = await fetch("http://localhost:8080/files/"+entryId);
        files = await resp.json();
        folders = files.filter(f => f.Size.Valid == false);
    });


</script>

<h1>{entry.Name}</h1>

<p>{files.length} files.</p>
<p>{folders.length} folders.</p>
<ul>
    {#each folders as f}
    {#if f.Filename.split('/').length==1}
    <li>{f.Filename}</li>
    {/if}
    {/each}
</ul>
<ul>
{#each scans as s}
<li>{s.Date} {s.Folder}</li>
{/each}
</ul>