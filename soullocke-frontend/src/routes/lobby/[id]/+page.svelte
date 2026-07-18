<script lang="ts">
	import type { PageProps } from './$types';
	import { getLobbyQuery } from '$api/lobby';
	import TextInput from '$components/textinput/TextInput.svelte';
	import PrimaryButton from '$components/buttons/PrimaryButton.svelte';
	import { getAccessQuery, getAuthMutation } from '$api/auth';

	let { params }: PageProps = $props();
	let password = $state('');

	const lobbyQuery = getLobbyQuery(() => params.id);
	const loginMutation = getAuthMutation();
	const accessCheck = getAccessQuery(() => params.id);

	const handleUnlockLobbyClick = (): void => {
		loginMutation.mutate({
			body: {
				lobbyId: params.id,
				password,
			},
		})
	};
</script>


{#if lobbyQuery.isLoading}
	Loading...
{/if}

{#if lobbyQuery.isError}
	{#if lobbyQuery.error.status === 404}
		<p>Lobby not found.</p>
	{:else}
		<p>An error occurred while fetching the lobby.</p>
	{/if}
{/if}

{#if lobbyQuery.isSuccess}
	<div class="mx-auto flex max-w-1/2 flex-col gap-4 mt-10">
		<div class="dark:bg-gray-800 rounded-2xl p-2">
			<h1 class="text-4xl text-center">{lobbyQuery.data.name}</h1>
		</div>
		{#if !(accessCheck.isSuccess && loginMutation.isSuccess)}
			<TextInput bind:value={password} label="Password" type="password"/>
			<PrimaryButton onClick={handleUnlockLobbyClick} variant="filled">Join Lobby</PrimaryButton>
		{/if}
		<p>Game Edition: {lobbyQuery.data.gameEditionId}</p>
		<p>WIP!</p>

		<ul>
			TODO
			<li>- Add encounters</li>
			<li>- Add players</li>
			<li>- Add utility tools (optional)</li>
			<li>- Some nice UI/UX</li>
		</ul>
	</div>
{/if}
