<script lang="ts">
    import Search from '$lib/components/Search.svelte';
    import ThFilter from '$lib/components/ThFilter.svelte';
    import ThSort from '$lib/components/ThSort.svelte';
    import RowCount from '$lib/components/RowCount.svelte';
    import RowsPerPage from '$lib/components/RowsPerPage.svelte';
    import Pagination from '$lib/components/Pagination.svelte';
    import { goto } from '$app/navigation';
    import { DataHandler } from '@vincjo/datatables/remote';
    import type { State, Row } from '@vincjo/datatables/remote';

    //import data from '$lib/data/api';

    import { reload } from '$lib/data/api'

    export let initialEquipments: Array<Row> = [];
    
    const handler = new DataHandler<Row>(initialEquipments, { rowsPerPage: 15, totalRows: 9999});
    const rows = handler.getRows();

    handler.onChange((state: State) => reload(state, "items"));

    function getImagePath(name: string) {
		const formattedName = name.toLowerCase().replace(/ /g, '_');
		return `/items/img/${formattedName}.png`;
    }

    // Function to navigate and include a query parameter
    function handleClick(index: string) {
        console.log('Button was clicked ' + index);
        // Navigate to a new page with a query parameter
        goto(`/item/${index}`);
    }

    const noImagePath = '/items/img/no_image.png';
    function handleError(event) {
        event.target.src = noImagePath;
    }
</script>

<div class=" overflow-x-auto space-y-2">
    <header class="flex justify-between gap-4">
        <Search {handler} />
        <RowsPerPage {handler} />
    </header>
    <table class="table table-hover table-compact table-auto w-full ">
        <thead>
            <tr>
                <!-- <ThSort {handler} orderBy="first_name">First name</ThSort>
                <ThSort {handler} orderBy="last_name">Last name</ThSort>
                <ThSort {handler} orderBy="email">Email</ThSort> -->
            </tr>
            <tr>
                <ThFilter {handler} filterBy="Name" />
                <ThFilter {handler} filterBy="Type" />
                <ThFilter {handler} filterBy="Base" />
            </tr>
        </thead>
        <tbody>
            {#each $rows as row, index}

                <tr on:click={() => handleClick(row.item_id)}>
                    <td><img src={getImagePath(row.Name)} on:error={handleError} alt={row.Name} class="image"/></td>
                    <td>{row.Rank}</td>
                    <td>{row.Name}</td>
                    <td>{row.Type}</td>
                    <td>{row.Base}</td>
                    <td>{row.Description}</td>
                    <td>{row.Size}</td>
                    <td>{row.Attributes}</td>
                </tr>

            {/each}
        </tbody>
    </table>
    <footer class="flex justify-between">
        <RowCount {handler} />
        <Pagination {handler} />
    </footer>
</div>