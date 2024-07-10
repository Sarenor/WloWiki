
<script>
    import { page } from '$app/stores'
    import data from '$lib/data/api';
    import dataAlchemy from '$lib/data/alchemy';
  //export let image = "/src/routes/items/imagespng/red_dragon_gown.png"; // Placeholder image URL

  // You can use this for additional interactivity or state management
  let isHovered = false;

  function getImagePath(name) {
		const formattedName = name.toLowerCase().replace(/ /g, '_');
		return `/items/img/${formattedName}.png`;
    }

  let item_id = $page.params.id;
  let item = data.find(item => item.item_id === parseInt(item_id));

  export let item_name = item?.Name;
  export let type = item?.Type;
  export let base = item?.Base
  export let description = item?.Description
  export let size = item?.Size
  export let rank = item?.Rank
  export let attributes = item?.Attributes;
</script>


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
<div class="card" on:mouseover={() => isHovered = true} on:mouseleave={() => isHovered = false}>
  <div class="card-header">
    <img src={getImagePath(item_name)} alt={item_name} />
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

<table class="skeleton">
  <thead>
    <tr>
      <th>Alchemy Level</th>
      <th>Recipe</th>
    </tr>
  </thead>
  <tbody>
    {#each dataAlchemy as { level, result, ingredients }}
      <tr>
        <td>{level}</td>
        <td>
          <div class="item">
            <img src={result.image} alt={result.name} />
            <span>{result.name}</span>
            =
            {#each ingredients as ingredient, index}
              <div class="item">
                <img src={ingredient.image} alt={ingredient.name} />
                <span>{ingredient.name}</span>
                {index < ingredients.length - 1 ? ' + ' : ''}
              </div>
            {/each}
          </div>
        </td>
      </tr>
    {/each}
  </tbody>
</table>


  <table class="skeleton">
    <thead>
      <tr>
        <th>Alchemy Level</th>
        <th>Recipe</th>
      </tr>
    </thead>
    <tbody>
      {#each dataAlchemy as { level, result, ingredients }}
        <tr>
          <td>{level}</td>
          <td>
            <div class="item">
              <img src={result.image} alt={result.name} />
              <span>{result.name}</span>
              =
              {#each ingredients as ingredient, index}
                <div class="item">
                  <img src={ingredient.image} alt={ingredient.name} />
                  <span>{ingredient.name}</span>
                  {index < ingredients.length - 1 ? ' + ' : ''}
                </div>
              {/each}
            </div>
          </td>
        </tr>
      {/each}
    </tbody>
  </table>