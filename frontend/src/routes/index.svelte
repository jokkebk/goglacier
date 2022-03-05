<script>
    import { onMount } from "svelte";

    let entries = [];

    onMount(async function() {
        let resp = await fetch('http://localhost:8080/entries');
        entries = await resp.json();
        let eMap = Object.fromEntries(entries.map(e => [e.Id, e]))
        
        resp = await fetch('http://localhost:8080/scans');
        const scans = await resp.json();
        scans.forEach(e => {
            const entry = eMap[e.EntryId];
            if(!('scans' in entry)) entry['scans'] = [];
            entry.scans.push(e);
        });
        entries = entries;
    });

    async function scan(e) {
        let folder = prompt('Enter folder to scan',
            e.scans ? e.scans.slice(-1)[0].Folder : '');
        
        let resp = await fetch('http://localhost:8080/scan', {
            method: 'POST', body: JSON.stringify({'EntryId': e.Id, folder})
        });
    }
</script>

<h1>GoGlacier</h1>

  {#each entries as e}
  <h2>{e.Name}</h2>
  {#if e.scans}
  <table>
      <thead>
          <tr>
              <th>Folder</th>
              <th>Date</th>
              <th>Skip</th>
              <th>Actions</th>
          </tr>
      </thead>
      <tbody>
    {#each e.scans as s}
        <tr>
            <td>{s.Folder}</td>
            <td>{s.Date}</td>
            <td>{s.Skip}</td>
            <td><a href="/scan/{s.Id}"><i class="bi-zoom-in"></i></a></td>
        </tr>
    {/each}
      </tbody>
  </table>
    {/if}
    
    <button on:click={scan(e)}>Scan now</button>
  {/each}

  <style>
      table {
          border-collapse: collapse;
      }

      th, td { border: 1px solid black; }
  </style>