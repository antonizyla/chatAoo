<script lang="ts">
	export let name: string = 'date';
	export let label: string = 'Date';

	export let minVal: Date = new Date(1900, 0, 1);
	export let maxVal: Date = new Date(2100, 0, 1);
	export let value: Date = new Date();
</script>

<div class="p-4">
	<label class="block py-2 pr-2" for={name}>
		{label}
	</label>
	<input
		class="rounded-md border-2 border-text/60 p-2 px-4"
		type="date"
		{name}
		{value}
		min={minVal.toString()}
		max={maxVal.toString()}
	/>
</div>
