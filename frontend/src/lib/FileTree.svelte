<script>
	import FileNode from "./FileNode.svelte";
	export let data; // type FileNode
</script>

<ul>
	<FileNode node={data} />	
</ul>

<style>
	ul {
		list-style-type: none;
		padding: 0;
		margin: 0;
	}
</style>