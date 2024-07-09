<script>
    // Define an array of objects to represent the table data
    // let tableArr = [
    //     { image: "src/routes/items/images/silvery_star.png", rank: "37 (76)", name: "Silvery Star", type: "Armor (Head)", base: "Corundum / Platinum", attributes: "MAT: +37" },
    //     { image: "src/routes/items/images/voodoo_doll.png", rank: "32 (64)", name: "Voodo Doll", type: "Armor (Head)", base: "Corundum / Platinum", attributes: "MAT: +37" },
    //     { image: "src/routes/items/images/jade_helmet.png", rank: "26 (52)", name: "Jade Helmet", type: "Armor (Head)", base: "Crystallization / Titanium / Magic", attributes: "SPD: +26" },
    //     // Add more rows as needed
    // ];
	import { onMount } from 'svelte';
	let items = [];
    let error = null;


	
    // Fetch items from the API
    async function fetchItems() {
        try {
            const response = await fetch('http://localhost:8080/api/items');
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            items = await response.json();
        } catch (err) {
            error = err.message;
        }
    }

    // Call fetchItems when the component is created
    onMount(fetchItems);

	    // Generate the image path based on the item name
	function getImagePath(name) {
		const formattedName = name.toLowerCase().replace(/ /g, '_').replace('-', '_');
		return `static/items/img/${formattedName}.png`;
    }
</script>

{#if error}
    <p>Error: {error}</p>
{/if}

<!-- Responsive Container (recommended) -->
<div class="table-container">
    <!-- Native Table Element -->
    <table class="table table-hover">
        <thead>
            <tr>
                <th>Image</th>
                <th>Rank</th>
                <th>Name</th>
                <th>Type</th>
				<th>Base</th>
				<th>Attributes</th>
            </tr>
        </thead>
        <tbody>
            {#each items as row}
                <tr>
					<td><img src={getImagePath(row.Name)} alt={row.name} class="image" /></td>
                    <td>{row.Rank}</td>
                    <td>{row.Name}</td>
                    <td>{row.Type}</td>
                    <td>{row.Base}</td>
					<td>{row.Attributes}</td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>

<style>
	.image {
        width: 50px; /* Adjust size as needed */
        height: auto; /* Maintain aspect ratio */
    }
    .table-container {
        overflow-x: auto;
        margin: 0 auto;
        max-width: 100%;
    }
    .table {
        width: 100%;
        border-collapse: collapse;
    }
    .table th, .table td {
        padding: 0.75em;
        border: 1px solid #3412cc;
    }
    .table th {
        background-color: #3b2391;
    }
    .table-hover tbody tr:hover {
        background-color: #f1f1f179;
    }
    tfoot {
        font-weight: bold;
    }
</style>
