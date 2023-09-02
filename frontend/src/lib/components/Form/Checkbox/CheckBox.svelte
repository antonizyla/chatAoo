<script lang="ts">
	export let label: string = '';
	export let checked: boolean = false;
</script>

<label class="container m-4 flex flex-row p-4 align-middle" for={label}>
	<input class="m-3" type="checkbox" name={label} value={label} bind:checked />
	<span class="checkmark" />
	<div class="m-3">
		{label}
	</div>
</label>

<style lang="postcss">
	input[type='checkbox'] {
		width: 1.5em;
		height: 1.5em;
		checkmark-color: #fff;
	}

	.checkmark:after {
		content: '';
		position: absolute;
		display: none;
	}

	.container input:checked ~ .checkmark {
		background-color: #aaa;
	}
</style>
