import { create, StateCreator, StoreMutatorIdentifier } from 'zustand';
import { createWithEqualityFn } from 'zustand/traditional';
import { devtools } from 'zustand/middleware';
import { immer } from 'zustand/middleware/immer';

export function createStore<T, Mos extends [StoreMutatorIdentifier, T][] = []>(
  store: StateCreator<T, [['zustand/immer', never]], Mos, T>,
  name: string = ''
) {
  const storeImmer = immer(store);
  if (process.env.NODE_ENV === 'development') {
    return create<T>()(
      devtools(storeImmer, {
        name: store.name || name,
        trace: true,
        store: store.name || name,
        anonymousActionType: 'setState',
      })
    );
  }
  return create<T>()(immer(store));
}

export function createStoreWithEqualityFn<T, Mos extends [StoreMutatorIdentifier, T][] = []>(
  store: StateCreator<T, [['zustand/immer', never]], Mos, T>,
  name: string = ''
) {
  if (process.env.NODE_ENV === 'development') {
    return createWithEqualityFn<T>()(
      devtools(immer(store), {
        name: store.name || name,
        trace: true,
        store: store.name || name,
        anonymousActionType: 'setState',
      }),
      Object.is
    );
  }
  return createWithEqualityFn<T>()(immer(store), Object.is);
}
