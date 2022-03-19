<script>
    import { slide } from 'svelte/transition';
	export let node;
	export let level = 0;

    const changes = 'chg new del'.split(' ');
	
	function toggle() {
		node.expanded = !node.expanded;
	}
</script>

<li on:click={toggle} style="padding-left:{level*1}rem" transition:slide>
    {#if !node.children || !node.children.length}
        <i class="bi-folder" /> 
	{:else if !node.expanded }
        <i class="bi-folder-plus" /> 
	{:else}
        <i class="bi-folder-minus" /> 
	{/if}
    {node.name}
    {#each changes as c}
    {#if node.count[c]}
    <span class={c}>{node.count[c]} {c}</span>
    {/if}
    {/each}
</li>

{#if node.expanded && node.children}
    {#each node.children as child}
        <svelte:self node={child} level={level+1}/>
    {/each}
{/if}

<style>
li {
    border-bottom: solid 1px #eee;
    margin: 0 0;
    padding: 0.2rem;
    background: #fafafa;
    //display: flex;
}

span.chg { background-color: #ccf; }
span.new { background-color: #cfc; }
span.del { background-color: #fcc; }
span.chg, span.new, span.del {
    border-radius: 0.5em;
    margin: 0.2em;
    padding: 0.2em;
    font-size: 80%;
}
</style>