import { reload } from '$lib/data/api';
import type { State } from '@vincjo/datatables/local';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	return {
        equipments: await reload({pageNumber: 0} as State, "items", fetch)
	};
};