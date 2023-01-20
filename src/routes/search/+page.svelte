<script lang="ts">
    export let data: any;
</script>

{#each data.cocktails as cocktail}
    <div>
        {cocktail.name}
    </div>
{/each}