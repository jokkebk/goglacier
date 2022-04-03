<script>
    import { onMount } from "svelte";

    let entries = {};

    onMount(async function () {
        let resp = await fetch("http://localhost:8080/scans");
        let scans = await resp.json();
        entries = scans.reduce((r,a) => {
            r[a.Entry] = r[a.Entry] || [];
            r[a.Entry].push(a);
            return r;
        }, Object.create(null));
    });

    async function getFiles(s) {
      if(s.Files) return;
      console.log(s.Date);
      let resp = await fetch("http://localhost:8080/filecount/"+s.Id);
      s.Files = await resp.json();
      entries = entries;
    }
    /*async function scan(e) {
        let folder = prompt(
            "Enter folder to scan",
            e.scans ? e.scans.slice(-1)[0].Folder : ""
        );

        let resp = await fetch("http://localhost:8080/scan", {
            method: "POST",
            body: JSON.stringify({ EntryId: e.Id, folder }),
        });
    }*/
</script>

<h1>GoGlacier</h1>

{#each Object.entries(entries) as [name, scans]}
<h2>{name}</h2>
<table>
    <thead>
        <tr>
            <th>Date</th>
            <th>Folder</th>
            <th>Skip</th>
            <th>Files</th>
            <th>Actions</th>
        </tr>
    </thead>
    <tbody>
        {#each scans as s}
            <tr>
                <td>{s.Date}</td>
                <td>{s.Folder}</td>
                <td>{s.Skip}</td>
                <!-- svelte-ignore a11y-mouse-events-have-key-events -->
                <td on:mouseover={getFiles(s)}>{s.Files || '?'}</td>
                <td><strong><a href="/scan/{s.Id}"> <i title="View" class="bi-zoom-in" /></a></strong></td>
            </tr>
        {/each}
    </tbody>
</table>
{/each}

<style>
    table {
        border-collapse: collapse;
    }

    th,
    td {
        border: 1px solid black;
        vertical-align: top;
    }
</style>
