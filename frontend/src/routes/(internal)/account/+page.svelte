<script lang="ts">
	import ChevronUp from '$lib/components/Accordion/chevron-up.svelte';
	import Button from '$lib/components/Button/Button.svelte';
	import { onMount } from 'svelte';

	let currentUser: string = '';
	let currentUserName: string = '';
	let displayName: string = '';
	let accountExists: boolean = true;
	onMount(() => {
		if (!localStorage.getItem('userID')) {
			accountExists = false;
		}
		currentUser = localStorage.getItem('userID') || '';
		currentUserName = localStorage.getItem('userName') || '';
		displayName = currentUserName;
	});

	async function changeUsername() {
		const update = await fetch(`http://localhost:8081/users/${currentUser}`, {
			method: 'PATCH',
			body: JSON.stringify({ name: currentUserName.trim() })
		});
		console.log(update);
		if (update.status === 200) {
			alert('Username Updated');
			localStorage.setItem('userName', currentUserName);
			displayName = currentUserName;
		} else {
			alert('Username Update Failed, Please Try Again');
		}
	}

	let validName: boolean = true;
	function validate() {
		if (currentUserName.length < 3) {
			validName = false;
		} else {
			validName = true;
		}
	}
</script>

<a href="/chat" class="p-2 block"><Button primary size="small">Home</Button></a>
{#if accountExists}
	<div class="w-fit mx-auto pt-40">
		<div class="p-2">
			<div class="p-2 text-center">
				You Are Currently Signed in as <div
					class="inline bg-primary text-white p-2 rounded-md mx-2.5"
				>
					{displayName}
				</div>
			</div>
			<div
				class="flex flex-row items-center gap-2 m-2 p-2 border-text border-solid border w-fit rounded-md"
			>
				<div class="flex flex-col gap-1.5">
					<input
						class="bg-primary-button/20 p-2 rounded-md"
						class:!bg-red-50={!validName}
						type="text"
						name="userName"
						id="userName"
						bind:value={currentUserName}
						on:input={validate}
						autofocus={true}
					/>
					{#if !validName}
						<p class="text-red-500 text-sm">Username must be at least 3 characters long</p>
					{/if}
				</div>
				<button on:click={changeUsername}><Button size="small">Change Username</Button></button>
			</div>
		</div>
	</div>
{:else}
	<p>You do not have an account stored in your browser</p>
	<a href="/"><Button>Create an Account</Button></a>
{/if}
