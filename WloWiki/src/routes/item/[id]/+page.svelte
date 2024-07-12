
<script>
  import { page } from '$app/stores'
  //import data from '$lib/data/api_static';
  //import dataAlchemy from '$lib/data/alchemy';
  import { onMount } from 'svelte';

  //export let image = "/src/routes/items/imagespng/red_dragon_gown.png"; // Placeholder image URL

  // You can use this for additional interactivity or state management
  let isHovered = false;
  let item = null;
  let recipes = []
  let error;

  function getImagePath(name) {
		const formattedName = name.toLowerCase().replace(/ /g, '_');
		return `/items/img/${formattedName}.png`;
    }

  let item_id = $page.params.id;
  // Fetch items from the API
  async function fetchItem() {
      try {
          const response = await fetch(`http://localhost:8080/api/items/${item_id}`);
          if (!response.ok) {
              throw new Error('Network response was not ok');
          }
          item = await response.json();
      } catch (err) {
          error = err;
      }
  }

  async function fetchRecipes() {
      try {
          const response = await fetch(`http://localhost:8080/alchemy?item_id=${item_id}`);
          if (!response.ok) {
              throw new Error('Network response was not ok');
          }
          recipes = await response.json();
      } catch (err) {
          error = err;
      }
  }
  

  onMount(async () => {
        await fetchItem();
        await fetchRecipes();
    });


  // Reactive statements to update exportable variables
  $: item_name = item?.Name;
  $: type = item?.Type;
  $: base = item?.Base;
  $: description = item?.Description;
  $: size = item?.Size;
  $: rank = item?.Rank;
  $: attributes = item?.Attributes;
  //$: recipe_data = recipes

</script>


<div class="card" on:mouseover={() => isHovered = true} on:mouseleave={() => isHovered = false}>
  <div class="card-header">
    {#if item_name}
        <img src={getImagePath(item_name)} alt={item_name} />
    {:else}
        <p>Loading...</p>
    {/if}
    <h2>{item_name}</h2>
  </div>
  <div class="card-content">
    <div class="info-row">
      <p><strong>Type:</strong> {type}</p>
      <p><strong>Rank:</strong> {rank}</p>
    </div>
    <div class="info-row">
      <p><strong>Attributes:</strong> {attributes}</p>
      <p><strong>Base:</strong> {base}</p>
    </div>
    <div class="info-row">
      <p><strong>Level:</strong> {rank}</p>
      <p><strong>Size:</strong> {size}</p>
    </div>
  </div>
  <div class="description">
    <p>{description}</p>
  </div>
  <div class="footer">
    {#if isHovered}
      <p>🧥 Check out this amazing armor!</p>
    {/if}
  </div>
</div>


{#if recipes.length === 0}
  <p>No recipes available.</p>
{:else}
  <table class="skeleton">
    <thead>
      <tr>
        <th>Alchemy Level</th>
        <th>Recipe</th>
      </tr>
    </thead>
    <tbody>
      {#each recipes as { level, rank_difference, result, ingredients }}
        <tr>
          <td>{level}</td>
          <td>
            <div class="item">
              <!-- <img src={result.image} alt={result.name} /> -->
              <span>{result}</span>
              =
              {#each ingredients as ingredient, index}
                <div class="item">
                  <!-- <img src={ingredient.image} alt={ingredient.name} /> -->
                  <span>{ingredient}</span>
                  {index < ingredients.length - 1 ? ' + ' : ''}
                </div>
              {/each}
            </div>
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
  {/if}

<style>
  .card {
    position: relative;
    border-radius: 12px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2); /* Increased shadow for better visibility */
    background-color: rgba(8, 154, 238, 0.1); /* Transparent background */
    border: 2px solid rgba(8, 154, 238, 0.5); /* Border with the same color as background */
    overflow: hidden;
    transition: transform 0.3s ease, box-shadow 0.3s ease;
    max-width: 450px; /* Set a max-width for consistency */
    margin: 0 auto; /* Center the card */
    padding: 1rem;
  }

  .card:hover {
    transform: scale(1.05); /* Slightly enlarge the card on hover */
    box-shadow: 0 6px 12px rgba(0, 0, 0, 0.2); /* Increase the shadow on hover */
  }

  .card-header {
    display: flex;
    align-items: center;
    gap: 1rem;
    margin-bottom: 1rem;
  }

  .card-header img {
    width: 50px; /* Set the width of the item image */
    height: 50px; /* Set the height of the item image */
    border-radius: 8px;
    object-fit: cover; /* Ensure image covers the box without distortion */
  }

  .card-header h2 {
    margin: 0;
  }

  .card-content {
    margin-bottom: 1rem;
  }

  .card-content .info-row {
    display: flex;
    justify-content: space-between;
    margin-bottom: 0.5rem;
  }

  .card-content p {
    margin: 0;
  }

  .footer {
    text-align: center;
    margin-top: 1rem;
    font-style: italic;
  }


    table {
    width: 100%;
    border-collapse: collapse;
  }

  th, td {
    padding: 10px;
    border: 1px solid #ddd;
    text-align: left;
  }

  .item {
    display: flex;
    align-items: center;
  }

  .item img {
    width: 32px;
    height: 32px;
    margin-right: 8px;
  }
  </style>


