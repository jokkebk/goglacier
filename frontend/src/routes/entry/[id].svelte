<script>
    import { page } from '$app/stores';
    import FileTree from '$lib/FileTree.svelte';
    import { onMount } from "svelte";

    let entry = {};
    let scans = [];
    let entryId = $page.params.id;
    let treeData = {};

    onMount(async function () {
        let resp = await fetch("http://localhost:8080/entries");
        let entries = await resp.json();
        entry = entries.find(e => e.Id == entryId);

        resp = await fetch("http://localhost:8080/scans");
        scans = await resp.json();
        scans = scans.filter(s => s.EntryId == entryId);

        resp = await fetch("http://localhost:8080/files/"+entryId);
        let files = await resp.json();
        let folders = files.filter(f => !f.Size.Valid);
        let folderMap = {};

        // We will see the folders in order, so children will come after parent
        for(let f of folders) {
            f.children = []; // initialize
            f.expanded = false;
            folderMap[f.Filename] = f; // save

            if(f.Filename == '.') {
                f.name = 'Root';
                f.expanded = true;
                continue; // we're done
            }

            let parts = f.Filename.split('/');
            f.name = parts.pop();
            let parentFolder = parts.length ? parts.join('/') : '.';
            f.parent = folderMap[parentFolder];
            f.parent.children.push(f);
        }
        treeData = folderMap['.'];
        console.log(treeData.children);
    });


</script>

<h1>{entry.Name}</h1>

<FileTree data={treeData}/>

<h2>Scans</h2>
<ul>
{#each scans as s}
<li>{s.Date} {s.Folder}</li>
{/each}
</ul>