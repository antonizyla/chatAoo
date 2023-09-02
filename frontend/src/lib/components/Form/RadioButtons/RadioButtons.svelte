<script lang="ts">
	export let labels: string[] = [];
	let bindingGroup = 1;
</script>

<div class="p-4">
	{#each labels as label}
		<label class="m-4" for={label}>
			<input class="m-2 ml-4" type="radio" name={label} value={label} bind:group={bindingGroup} />
			{label}
		</label>
	{/each}
</div>
