<script>
    import { onMount } from "svelte";

    let entries = [];

    onMount(async function () {
        let resp = await fetch("http://localhost:8080/entries");
        entries = await resp.json();
        let eMap = Object.fromEntries(entries.map((e) => [e.Id, e]));

        resp = await fetch("http://localhost:8080/scans");
        const scans = await resp.json();
        scans.forEach((e) => {
            const entry = eMap[e.EntryId];
            if (!("scans" in entry)) entry["scans"] = [];
            entry.scans.push(e);
        });
        entries = entries;
    });

    async function scan(e) {
        let folder = prompt(
            "Enter folder to scan",
            e.scans ? e.scans.slice(-1)[0].Folder : ""
        );

        let resp = await fetch("http://localhost:8080/scan", {
            method: "POST",
            body: JSON.stringify({ EntryId: e.Id, folder }),
        });
    }
</script>

<h1>GoGlacier</h1>

<table>
    <thead>
        <tr>
            <th>Entry</th>
            <th>Scans</th>
            <th>Actions</th>
        </tr>
    </thead>
    <tbody>
        {#each entries as e}
            <tr>
                <td>{e.Name}</td>
                <td>
                    {#if e.scans}
                    <p>
                    {#if e.scans.length > 3}
                        ...<br>
                    {/if}
                        {#each e.scans.slice(-3) as s}
                            {s.Date}<br>
                        {/each}
                    </p>
                    {/if}
                </td>
                <td><a href="/entry/{e.Id}">
                        <i class="bi-zoom-in" /> View entry</a>
                    </td>
            </tr>
        {/each}
    </tbody>
</table>

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
