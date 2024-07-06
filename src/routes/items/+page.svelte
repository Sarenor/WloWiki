<script>
	import {
	  Table,
	  TableBody,
	  TableBodyCell,
	  TableBodyRow,
	  TableHead,
	  TableHeadCell,
	  ImagePlaceholder,
	  Modal
	} from 'flowbite-svelte';
	import { slide } from 'svelte/transition';
	import welcome from '$lib/images/silvery_star.webp';
	import welcome_fallback from '$lib/images/silvery_star.png';
	const items = [
	  {
		rank: '37 (76)',
		name: 'Silvery Star',
		type: "Armor (Head)",
		base: "Corundum / Platinum",
		attributes: "MAT:+37",
	  },
	  {
		rank: '32 (64)',
		name: 'Voodoo Doll',
		type: "Armor (Head)",
		base: "Magic / Grass / Corundum",
		attributes: "MAT:+24 ",
	  },
	  {
		rank: '36 (72)',
		name: 'Eagle Headwear',
		type: "Armor (Head)",
		base: "Steel / Pure Iron / Corundum",
		attributes: "ATK:+22 SPD:+14",
	  },
	];
  
	let openRow
	let details
	let doubleClickModal = false
  
	const toggleRow = (i) => {
	  openRow = openRow === i ? null : i
	}
  
  </script>
  
  <Table>
	<TableHead>
		<TableHeadCell>Rank</TableHeadCell>
	  <TableHeadCell>Name</TableHeadCell>
	  <TableHeadCell>Type</TableHeadCell>
	  <TableHeadCell>Base</TableHeadCell>
	  <TableHeadCell>Attributes</TableHeadCell>
	</TableHead>
	<TableBody tableBodyClass="divide-y">
	  {#each items as item, i}
		<TableBodyRow on:click={() => toggleRow(i)}>
		<TableBodyCell>				<picture>
			<source srcset={welcome} type="image/webp" />
			<img src={welcome_fallback} alt="Welcome" />
		</picture></TableBodyCell>
		  <TableBodyCell>{item.rank}</TableBodyCell>
		  <TableBodyCell>{item.name}</TableBodyCell>
		  <TableBodyCell>{item.type}</TableBodyCell>
		  <TableBodyCell>{item.base}</TableBodyCell>
		  <TableBodyCell>{item.attributes}</TableBodyCell>
		</TableBodyRow>
		{#if openRow === i}
		  <TableBodyRow on:dblclick={() => {
			doubleClickModal = true;
			details = item;
		  }}>
			<TableBodyCell colspan="4" class="p-0">
			  <div class="px-2 py-3" transition:slide={{ duration: 300, axis: 'y' }}>
				<!--<ImagePlaceholder />-->

			  </div>
			</TableBodyCell>
		  </TableBodyRow>
		{/if}
	  {/each}
	</TableBody>
  </Table>
  <Modal title={details?.name} bind:open={doubleClickModal} autoclose outsideclose>
	<ImagePlaceholder />
  </Modal>