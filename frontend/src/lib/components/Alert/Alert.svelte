<script lang="ts">
	export let dialog: HTMLDialogElement;
</script>

<dialog class="rounded-lg" bind:this={dialog} on:click|self={() => dialog.close()}>
	<div on:click|stopPropagation>
		<slot />
	</div>
</dialog>
