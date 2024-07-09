<script>
    let name = '';
    let rank = '';
    let selectedBases = [];
    let type = '';
    let attributes = '';
    let error = '';

    // Handle form submission
    async function handleSubmit() {
        const itemData = {
            name,
            rank,
            base: selectedBases.join(", "),
            type,
            attributes
        };

        try {
            const response = await fetch('http://localhost:8080/api/items', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(itemData)
            });

            if (response.ok) {
                console.log('Item added successfully');
                // Reset the form
                name = '';
                rank = '';
                selectedBases = [];
                type = '';
                attributes = '';
                error = '';
            } else {
                error = `Failed to add item: ${response.statusText}`;
            }
        } catch (err) {
            error = `Error: ${err.message}`;
        }
    }
</script>

<form on:submit|preventDefault={handleSubmit}>
    <label class="label">
        <span>Item Name</span>
        <input class="input" type="text" name="name" placeholder="" bind:value={name} />
    </label>

    <label class="label">
        <span>Rank</span>
        <input class="input" type="text" name="rank" placeholder="" bind:value={rank} />
    </label>

    <label class="label">
        <span>Base</span>
        <div class="space-y-2">
            <label class="flex items-center space-x-2">
                <input class="checkbox" type="checkbox" value="Platinum" bind:group={selectedBases} />
                <p>Platinum</p>
            </label>
            <label class="flex items-center space-x-2">
                <input class="checkbox" type="checkbox" value="Crystallization" bind:group={selectedBases} />
                <p>Crystallization</p>
            </label>
            <label class="flex items-center space-x-2">
                <input class="checkbox" type="checkbox" value="Titanium" bind:group={selectedBases} />
                <p>Titanium</p>
            </label>
            <label class="flex items-center space-x-2">
                <input class="checkbox" type="checkbox" value="Magic" bind:group={selectedBases} />
                <p>Magic</p>
            </label>
            <label class="flex items-center space-x-2">
                <input class="checkbox" type="checkbox" value="Corundum" bind:group={selectedBases} />
                <p>Corundum</p>
            </label>
        </div>
    </label>

    <label class="label">
        <span>Type</span>
        <select class="select" name="type" bind:value={type}>
            <option value="Armor (Head)">Armor (Head)</option>
            <option value="Armor (Gloves)">Armor (Gloves)</option>
            <option value="Armor (Chest)">Armor (Chest)</option>
            <option value="Armor (Boots)">Armor (Boots)</option>
            <option value="Armor (Special)">Armor (Special)</option>
        </select>
    </label>

    <label class="label">
        <span>Attributes</span>
        <input class="input" type="text" name="attributes" placeholder="" bind:value={attributes} />
    </label>

    <button type="submit" class="btn variant-filled">Add</button>

    {#if error}
        <p class="error">{error}</p>
    {/if}
</form>