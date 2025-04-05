import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Search } from "lucide-react";

export default function SearchCard() {
  return (
    <Card className="w-full">
      <CardHeader>
        <CardTitle>Rechercher des données compromises</CardTitle>
        <CardDescription>Entrez votre adresse e-mail pour vérifier si elle apparaît dans une violation de données</CardDescription>
      </CardHeader>
      <CardContent>
        <form id="search-section">
          <div className="flex w-full max-w-sm items-center space-x-2">
            <Input type="email" placeholder="Entrez votre e-mail" />
            <Button type="submit">
              <Search className="h-4 w-4 mr-2" />
              Rechercher
            </Button>
          </div>
        </form>
      </CardContent>
    </Card>
  );
} 