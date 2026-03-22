import { createMutation, createQuery } from '@tanstack/svelte-query';
import { createLobbyMutation, getLobbyOptions } from './generated/@tanstack/svelte-query.gen.ts';

export function getLobbyQuery(getId: () => string) {
	return createQuery(() => {
		const id = getId();

		return {
			...getLobbyOptions({ path: { id } }),
			enabled: id.length > 0
		};
	});
}

export const getLobbyCreationMutation = () => {
	return createMutation(() => {
		return {
			...createLobbyMutation()
		}
	})
};
