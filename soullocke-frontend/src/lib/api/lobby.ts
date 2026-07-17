import { createMutation, createQuery } from '@tanstack/svelte-query';
import { createLobbyMutation, getLobbyOptions } from './generated/@tanstack/svelte-query.gen';

export function getLobbyQuery(getId: () => string) {
	return createQuery(() => {
		const id = getId();

		return {
			...getLobbyOptions({ path: { lobbyId: id } }),
			enabled: id.length > 0
		};
	});
}

export const getLobbyCreationMutation = () => {
	return createMutation(() => {
		return {
			...createLobbyMutation(),
		}
	})
};
