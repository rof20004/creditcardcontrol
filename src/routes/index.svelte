<style>
	h1, figure, p {
		text-align: center;
		margin: 0 auto;
	}

	h1 {
		font-size: 2.8em;
		text-transform: uppercase;
		font-weight: 700;
		margin: 0 0 0.5em 0;
	}

	h2 {
		text-align: center;
		font-size: 1.8em;
		margin-bottom: 50px;
	}

	figure {
		margin: 0 auto;
	}

	img {
		width: 100%;
		max-width: 400px;
		margin: 0 auto;
		margin-bottom: 50px;
	}

	p {
		margin: 1em auto;
	}

	@media (min-width: 480px) {
		h1 {
			font-size: 4em;
		}
	}
</style>

<svelte:head>
	<title>Sapper project template</title>
</svelte:head>

<script>
async function createCard(method = 'POST') {
  try {
    const response = await fetch('/.netlify/functions/create_card', { method: method })
    const data = await response.json()
    console.log(data)
  } catch (error) {
    console.log(error)
  }
}

async function deleteCard(method = 'DELETE') {
  try {
    const response = await fetch('/.netlify/functions/delete_card', { method: method })
    const data = await response.json()
    console.log(data)
  } catch (error) {
    console.log(error)
  }
}
</script>

<h1>Great success!</h1>

<h2><u>Svelte + Sapper + Tailwindcss + PurgeCSS</u></h2>

<button on:click="{ createCard }">Post lambda function</button>
<button on:click="{ () => createCard('GET') }">Post lambda function with error</button><br>

<button on:click="{ deleteCard }">Delete lambda function</button>
<button on:click="{ () => deleteCard('GET') }">Delete lambda function with error</button><br>

<figure>
	<img alt='Borat' src='great-success.png'>
	<figcaption>HIGH FIVE!</figcaption>
</figure>

<p><strong>Try editing this file (src/routes/index.svelte) to test live reloading.</strong></p>
