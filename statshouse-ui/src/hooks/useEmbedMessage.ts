import { useRectObserver } from './useRectObserver';
import { useEffect } from 'react';

export function useEmbedMessage(refPage: Element | null, embed?: boolean) {
  const [{ width, height }] = useRectObserver(refPage, true);
  useEffect(() => {
    if (embed) {
      window.top?.postMessage({ source: 'statshouse', payload: { width, height } }, '*');
    }
  }, [embed, height, width]);
}
