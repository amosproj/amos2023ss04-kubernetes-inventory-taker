"use client";

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function H1({ content }: { content: any }): JSX.Element {
  return (
    <h1 className="mb-4 my-2 text-3xl font-bold leading-none tracking-tight text-gray-900 dark:text-white">
      {content}
    </h1>
  );
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function H2({ content }: { content: any }): JSX.Element {
  return (
    <h2 className="mb-4 my-2 text-3xl font-bold leading-none tracking-tight text-gray-900 dark:text-white">
      {content}
    </h2>
  );
}
