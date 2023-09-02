<script lang="ts">
	export let name: string;
	export let multiple: boolean = false;
</script>

<input class="w-full rounded-md border border-red-50 p-4" type="file" {name} {multiple} />

<style lang="postcss">
	input[type='file']::file-selector-button {
		@apply mr-10 inline-block scale-105 cursor-pointer rounded-md border-2 border-secondary-button bg-secondary-button px-6 py-3 shadow-2xl shadow-secondary-button transition duration-200 ease-in;
	}
</style>
