<script lang="ts">
	import Button from '../Button/Button.svelte';

	let open = false;
</script>

<Button
	on:click={() => {
		open = !open;
	}}>Open</Button
>
<div class:hidden={!open} class:block={open}>
	<Button>a</Button>
	<Button>b</Button>
</div>
