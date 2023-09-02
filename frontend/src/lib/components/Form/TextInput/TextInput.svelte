<script lang="ts">
	export let placeholder = 'Your Name...';
	export let id = 'email';
	export let label: string;
	export let required: boolean = false;
	export let validate = true;
	export let validationMessage = '* Required';
	export let autocomplete = false;
	export let name: string;
	export let regex: string = '(.)';

	const re = new RegExp(regex);

	let input = '';
	let valid = true;

	function validateMessage(message: string) {
		if (validate && message.length > 0) {
			valid = re.test(message);
		} else {
			valid = true;
		}
	}

	$: validateMessage(input);
</script>

<div class="flex w-full flex-col p-2">
	<div class="flex flex-row">
		<div class="p-1">{label}</div>
		{#if required}
			<div class="p-1 text-center text-sm text-red-400">*</div>
		{/if}
	</div>
	<input
		bind:value={input}
		class="min-w-full rounded-md border-2 border-text/40 p-3 shadow-sm"
		class:border-red-400={!valid}
		type="text"
		{name}
		{id}
		{placeholder}
		{required}
		autocomplete={autocomplete ? 'on' : 'off'}
	/>
	{#if !valid}
		<div class="p-1 text-center text-sm text-red-400">{validationMessage}</div>
	{/if}
</div>
