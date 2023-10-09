<script lang="ts">
	import { onMount } from 'svelte';
	import Button from '$lib/components/Button/Button.svelte';

	let userName: string | null = '';
	let userID: string | null = '';
	let mounted = false;
	onMount(async () => {
		// check local storage for a user tag
		mounted = true;
		userName = localStorage.getItem('userName');
		userID = localStorage.getItem('userID');
	});
</script>

<div class="p-12 w-fit mx-auto mt-20">
	{#if userName}
		<div class="flex flex-col items-center gap-2">
			<div class="p-3">
				You are currently signed in as <p
					class="bg-primary text-background p-2 rounded-md inline mx-2"
				>
					{userName}
				</p>
			</div>
			<div class="flex flex-row p-2 gap-2">
				<a href="/account"><Button>Edit Account Details</Button></a>
				<a href="/chat"><Button primary>Create or Enter a chat</Button></a>
			</div>
		</div>
	{:else}
		<p class="p-2">No User Identifier has been found in your browser</p>
		<a href="/account/create">
			<Button>Create an Account</Button>
		</a>
	{/if}
</div>
