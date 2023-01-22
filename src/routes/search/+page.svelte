<script lang="ts">
	let cocktails: any[] = [];

	const load = async () => {
		const res = await fetch(`/api/cocktails-thebar.json`);
		cocktails = await res.json();

		return { cocktails };
	};
</script>

{#each cocktails as cocktail}
	<div>
		{cocktail.name}
	</div>
{/each}
