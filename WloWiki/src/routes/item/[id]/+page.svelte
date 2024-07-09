
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
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
      background-color: #ffffff;
      overflow: hidden;
      transition: transform 0.3s ease, box-shadow 0.3s ease;
      max-width: 300px; /* Set a max-width for consistency */
      margin: 0 auto; /* Center the card */
    }
  
    .card:hover {
      transform: translateY(-10px);
      box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
    }
  
    .card-header {
      position: relative;
      text-align: center;
      display: flex;
      justify-content: center;
      align-items: center;
      height: 50px; /* Set a fixed height for the header */
    }
  
    .card-header img {
      max-width: 100%;
      max-height: 100%;
      width: auto;
      height: auto;
      object-fit: contain;
      transform: scale(1.5); /* Scale up the image */

    }
  
    .card-content {
      padding: 16px;
      text-align: center;
    }
  
    .card-content h2 {
      margin: 0;
      font-size: 1.5rem;
      font-weight: bold;
      color: #333;
    }
  
    .card-content p {
      margin: 8px 0;
      color: #555;
      font-size: 0.9rem;
    }
  
    .card-stats {
      margin-top: 12px;
      text-align: center;
    }
  
    .card-stats .stat {
      background: #f0f4f8;
      border-radius: 8px;
      padding: 8px;
      margin: 4px 0;
      font-size: 0.9rem;
      color: #666;
    }
  
    .footer {
      position: absolute;
      bottom: 16px;
      left: 16px;
      right: 16px;
      text-align: center;
      font-size: 0.8rem;
      color: #888;
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
  
  <div class="card p-4 custom-width" on:mouseover={() => isHovered = true} on:mouseleave={() => isHovered = false}>
    <div class="card-header">
      <td><img src={getImagePath(item_name)} alt="{item_name}" /></td>
    </div>
    <div class="card-content">
      <h2>{item_name}</h2>
      <p><strong>Type:</strong> {type}</p>
      <p><strong>Base:</strong> {base}</p>
      <p><strong>Size:</strong> {size}</p>
      <p><strong>Rank:</strong> {rank}</p>
      <p><strong>Attributes:</strong> {attributes}</p>
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