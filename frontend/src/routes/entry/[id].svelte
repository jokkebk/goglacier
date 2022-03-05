<script>
    import { page } from '$app/stores';
    import { onMount } from "svelte";

    let entry = [];
    let scans = [];
    let entryId = $page.params.id;

    onMount(async function () {
        let resp = await fetch("http://localhost:8080/entries");
        let entries = await resp.json();
        entry = entries.find(e => e.Id == entryId);

        resp = await fetch("http://localhost:8080/scans");
        scans = await resp.json();
        scans = scans.filter(s => s.EntryId == entryId);
    });


</script>

<h1>{entry.Name}</h1>

<ul>
{#each scans as s}
<li>{s.Date} {s.Folder}</li>
{/each}
</ul>