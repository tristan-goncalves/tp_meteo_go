# Partie A - Exploration des deux fichiers

## Différences observées entre les deux formats

| Donnée | Représentation JSON                                                         | Représentation XML                                                       |
|---|-----------------------------------------------------------------------------|--------------------------------------------------------------------------|
| Pays | Le champ `"country": "France"` est avec le nom complet du pays              | L'attribut est `country="FR"` et est directement dans `<station>`        |
| Coordonnées | L'objet `"location"` est détaillé avec `"latitude"` et `"longitude"`        | On a un élément `<coordinates>` avec des attributs `lat` et `lon`        |
| Altitude | Le champ `"altitude_m": 47` est dans la station                             | L'attribut `altitude="47"` est dans `<coordinates>`                      |
| Modèle de capteur | On a un objet `"device"` avec `"type"`, `"manufacturer"` et `"installed_on"` | On a un élément `<hardware>` avec attributs `model`, `vendor` et `since` |
| Température | On a un champ `"temperature_celsius": 3.3` dans chaque observation          | On a un élément `<measure type="temperature" unit="C">3.3</measure>`     |
| Conditions météo | On a un champ `"conditions": "clear"`                                       | On a un attribut `sky="clear"` sur `<observation>`                       |
| Vent | Objet `"wind"` avec `"speed_kmh"` et `"direction_deg"`                      | Élément `<wind>` avec attributs `speed` et `direction`                   |
| Notes optionnelles | On a un champ `"notes": null`                        | L'élément `<note>` n'est présent uniquement si une note existe           |
