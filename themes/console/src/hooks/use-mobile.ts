/**
 * 移动端检测 Hook
 */

import { useState, useEffect } from "react";

const MOBILE_BREAKPOINT = 768;

export function useIsMobile(): boolean {
  const [isMobile, setIsMobile] = useState(false);

  useEffect(() => {
    const mql = window.matchMedia(`(max-width: ${MOBILE_BREAKPOINT - 1}px)`);

    const onChange = () => setIsMobile(window.innerWidth < MOBILE_BREAKPOINT);

    mql.addEventListener("change", onChange);
    onChange();

    return () => mql.removeEventListener("change", onChange);
  }, []);

  return isMobile;
}
