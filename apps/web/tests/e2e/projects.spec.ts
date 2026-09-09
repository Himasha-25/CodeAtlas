import { test, expect } from "@playwright/test";

test.describe("Projects", () => {
  test.beforeEach(async ({ page }) => {
    // Assumes a valid token is already in localStorage via a setup fixture.
    await page.goto("/projects");
  });

  test("shows empty state", async ({ page }) => {
    await expect(page.getByText("No projects yet.")).toBeVisible();
  });

  test("creates a project", async ({ page }) => {
    await page.getByRole("button", { name: "New project" }).click();
    await page.fill('input[id="name"]', "My Project");
    await page.getByRole("button", { name: "Create" }).click();
    await expect(page.getByText("My Project")).toBeVisible();
  });
});
