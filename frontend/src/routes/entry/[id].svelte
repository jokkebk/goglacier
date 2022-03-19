<script>
    import { page } from '$app/stores';
    import FileTree from '$lib/FileTree.svelte';
    import { onMount } from "svelte";

    let entryId = $page.params.id;
    let entry = {};

    let scans = [];
    let selected; // selected scan

    let files = [];
    let treeData;


    $: if(selected && selected.Id) {
        let base = {};
        let folderMap = {'': { name: 'Root', expanded: true, children: [],
                        count: {'chg': 0, 'new': 0, 'del': 0} }};

        // We will see the folders in order, so children will come after parent
        for(let f of files) {
            if(f.ScanId < selected.Id) { // build base
                if(f.Modified !== null) base[f.Filename] = f;
                else delete base[f.Filename];
            } else if(f.ScanId == selected.Id) { // build deltas
                let change, type;
                if(f.Modified === null) {
                    type = base[f.Filename].Size === null ? 'dir' : 'file';
                    change = 'del';
                } else {
                    type = f.Size === null ? 'dir' : 'file';
                    change = (f.Filename in base) ? 'chg' : 'new';
                }
                if(change=='chg' && type=='dir') continue; // don't care
                //console.log(change, type, f.Filename);
                
                // Propagate delta to folder tree
                let parts = f.Filename.split('/');
                for(let i=0; i<=parts.length - (type=='file'?1:0); i++) {
                    let path = parts.slice(0,i).join('/'), dir;
                    if(path in folderMap) dir = folderMap[path];
                    else folderMap[path] = dir = { // never i=0, ''
                        name: parts[i-1], expanded: false, children: [],
                        count: {'chg': 0, 'new': 0, 'del': 0}
                    };
                    dir.count[change]++;
                    //console.log(path, dir);
                }
            } else break; // end after selected id
        }

        // For easier tracking, add top level folders from base
        for(const key of Object.keys(base)) {
            const f = base[key];
            if(f.Size !== null || f.Filename.indexOf('/')!=-1 ||
                f.Filename in folderMap) continue;
            folderMap[f.Filename] = {name: f.Filename,
                children: [], count: {'chg':0,'new':0,'del':0}};
        }

        // Add deltas
        for(const key of Object.keys(folderMap).sort()) {
            if(key === '') continue;
            const dir = folderMap[key];
            const parentDir = key.split('/').slice(0,-1).join('/');
            const parent = folderMap[parentDir];
            //console.log(`Adding ${dir.name} to ${parent.name}...`);
            parent.children.push(dir);
        }
        
        treeData = folderMap[''];
    }

    onMount(async function () {
        let resp = await fetch("http://localhost:8080/entries");
        let entries = await resp.json();
        entry = entries.find(e => e.Id == entryId);

        resp = await fetch("http://localhost:8080/scans");
        scans = await resp.json();
        scans = scans.filter(s => s.EntryId == entryId);

        resp = await fetch("http://localhost:8080/files/"+entryId);
        files = await resp.json();
        
        selected = scans.slice(-1)[0];
    });
</script>

<h1>{entry.Name}</h1>

<select bind:value={selected}>
{#each scans as s}
<option value={s}>#{s.Id}: {s.Date} {s.Folder}</option>
{/each}
</select>

{#if treeData}
<FileTree data={treeData}/>
{/if}

<h2>Scans</h2>
<ul>
</ul>