<script>
    import MaterialDatatable from './MaterialDatatable.svelte';
	export let data;
	console.log(data)

	
</script>

<MaterialDatatable initialMaterials={data.materials} />