import { describe, it, expect } from "vitest";
import { render, screen } from "@testing-library/react";
import { Badge } from "@/components/ui/badge";

describe("Badge", () => {
  it("renders text", () => {
    render(<Badge>completed</Badge>);
    expect(screen.getByText("completed")).toBeDefined();
  });

  it("applies variant class", () => {
    const { container } = render(<Badge variant="success">ok</Badge>);
    expect(container.firstChild).toHaveProperty("className");
  });
});
