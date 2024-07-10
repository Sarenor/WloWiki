<script>
    let pages = [
      { name: 'Items', path: '/equipment' },
      { name: 'Alchemy', path: '/alchemy' },
      { name: 'Monsters', path: '/monsters' },
      { name: 'Manufacture', path: '/manufacture' },
      { name: 'Quests', path: '/quests' },
      { name: 'Guides', path: '/guides' },
      { name: 'About', path: '/about' }
    ];
  </script>
  
  <div class="navbar">
    <ul>
      {#each pages as page}
        <li><a href={page.path}>{page.name}</a></li>
      {/each}
    </ul>
  </div>
  
  <div class="main-content">
    <h1>Welcome to the Witch's Domain</h1>
    <img src="maid.png" alt="Glowing Witch" class="witch-image" />
  
    <div class="change-fixes">
      <h2>Change Fixes</h2>
      <p>Add your change fixes here...</p>
    </div>
  </div>
  
  <style>
    .navbar {
      background-color: #100b52;
      overflow: hidden;
      position: fixed;
      width: 100%;
      top: 0;
      z-index: 1000;
    }
  
    .navbar ul {
      list-style-type: none;
      margin: 0;
      padding: 0;
      display: flex;
      justify-content: center;
    }
  
    .navbar li {
      margin: 0 10px;
    }
  
    .navbar a {
      display: block;
      color: white;
      text-align: center;
      padding: 14px 16px;
      text-decoration: none;
    }
  
    .navbar a:hover {
      background-color: #ddd;
      color: black;
    }
  
    .main-content {
      text-align: center;
      padding: 80px 20px 20px; /* Added top padding to account for the fixed navbar */
    }
  
    .witch-image {
      width: 300px;
      height: auto;
      margin: 0 auto;
      display: block;
      animation: glow 1.5s infinite alternate;
    }
  
    @keyframes glow {
      from {
        box-shadow: 0 0 10px #ffd700;
      }
      to {
        box-shadow: 0 0 20px #ffd700;
      }
    }
  
    .change-fixes {
      margin-top: 40px;
      padding: 20px;
      border: 1px solid #ccc;
      border-radius: 8px;
      background-color: #f9f9f9; /* Light background for contrast */
      display: inline-block; /* Centered with respect to the image */
      box-shadow: 0 0 10px rgba(0, 0, 0, 0.1); /* Subtle shadow for better visibility */
    }
  
    .change-fixes h2 {
      margin-top: 0;
    }
  </style>
  