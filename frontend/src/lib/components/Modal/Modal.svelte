<script lang="ts">
	export let showing: boolean;

	let dialog: HTMLDialogElement;
	$: if (dialog && showing) dialog.showModal();
</script>

<dialog
	class="max-w-3xl rounded-md border-none p-0 backdrop:bg-black/30"
	bind:this={dialog}
	on:close={() => {
		showing = false;
	}}
	on:click|self={() => dialog.close()}
>
	<div class="" on:click|stopPropagation>
		<div class="flex-col flex p-2 w-80 max-w-[95vw]">
			<button class="text-right px-1" autofocus on:click={() => dialog.close()}>X</button>
			<slot />
		</div>
	</div>
</dialog>
